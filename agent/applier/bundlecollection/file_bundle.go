package bundlecollection

import (
	"os"
	"path"
	"path/filepath"

	"code.cloudfoundry.org/clock"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-utils/fileutil"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
)

const (
	fileBundleLogTag = "FileBundle"
)

type FileBundle struct {
	installPath  string
	enablePath   string
	fileMode     os.FileMode
	fs           boshsys.FileSystem
	timeProvider clock.Clock
	compressor   fileutil.Compressor
	logger       boshlog.Logger
}

func NewFileBundle(
	installPath, enablePath string,
	fileMode os.FileMode,
	fs boshsys.FileSystem,
	timeProvider clock.Clock,
	compressor fileutil.Compressor,
	logger boshlog.Logger,
) FileBundle {
	return FileBundle{
		installPath:  installPath,
		enablePath:   enablePath,
		fileMode:     fileMode,
		fs:           fs,
		timeProvider: timeProvider,
		compressor:   compressor,
		logger:       logger,
	}
}

func (b FileBundle) InstallWithoutContents() (string, error) {
	b.logger.Debug(fileBundleLogTag, "Installing without contents %v", b)

	if err := b.fs.MkdirAll(b.installPath, b.fileMode); err != nil {
		return "", bosherr.WrapError(err, "Creating parent installation directory")
	}
	if err := b.fs.Chown(path.Dir(b.installPath), "root:vcap"); err != nil {
		_ = b.Uninstall()
		return "", bosherr.WrapError(err, "Setting ownership on parent installation directory")
	}
	if err := b.fs.Chown(b.installPath, "root:vcap"); err != nil {
		_ = b.Uninstall()
		return "", bosherr.WrapError(err, "Setting ownership on installation directory")
	}

	return b.installPath, nil
}

func (b FileBundle) Install(sourcePath, pathInBundle string) (string, error) {
	b.logger.Debug(fileBundleLogTag, "Installing %v", b)

	if _, err := b.InstallWithoutContents(); err != nil {
		return "", err
	}

	// Job bundles contain more than one job. We receive the individual job's
	// path as pathInBundle but we don't want to have that be duplicated in the
	// installPath so we have tar strip that component out.
	stripComponents := 0
	if pathInBundle != "" {
		stripComponents = 1
	}

	err := b.compressor.DecompressFileToDir(
		sourcePath,
		b.installPath,
		fileutil.CompressorOptions{PathInArchive: pathInBundle, StripComponents: stripComponents},
	)
	if err != nil {
		if pathInBundle != "" {
			err = b.compressor.DecompressFileToDir(
				sourcePath,
				b.installPath,
				fileutil.CompressorOptions{PathInArchive: "./" + pathInBundle, StripComponents: stripComponents},
			)
		}
		if err != nil {
			_ = b.Uninstall()
			return "", bosherr.WrapError(err, "Decompressing package files")
		}
	}

	b.logger.Debug(fileBundleLogTag, "Installing %v", b)
	return b.installPath, nil
}

func (b FileBundle) GetInstallPath() (string, error) {
	path := b.installPath
	if !b.fs.FileExists(path) {
		return "", bosherr.Error("install dir does not exist")
	}

	return path, nil
}

func (b FileBundle) IsInstalled() (bool, error) {
	return b.fs.FileExists(b.installPath), nil
}

func (b FileBundle) Enable() (string, error) {
	b.logger.Debug(fileBundleLogTag, "Enabling %v", b)

	if !b.fs.FileExists(b.installPath) {
		return "", bosherr.Error("bundle must be installed")
	}

	err := b.fs.MkdirAll(filepath.Dir(b.enablePath), b.fileMode)
	if err != nil {
		return "", bosherr.WrapError(err, "failed to create enable dir")
	}

	err = b.fs.Chown(filepath.Dir(b.enablePath), "root:vcap")
	if err != nil {
		return "", bosherr.WrapError(err, "Setting ownership on source directory")
	}

	err = b.fs.Symlink(b.installPath, b.enablePath)
	if err != nil {
		return "", bosherr.WrapError(err, "failed to enable")
	}

	return b.enablePath, nil
}

func (b FileBundle) Disable() error {
	b.logger.Debug(fileBundleLogTag, "Disabling %v", b)

	target, err := b.fs.ReadAndFollowLink(b.enablePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return bosherr.WrapError(err, "Reading symlink")
	}

	if target == b.installPath {
		return b.fs.RemoveAll(b.enablePath)
	}

	return nil
}
