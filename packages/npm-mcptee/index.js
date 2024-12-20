const path = require("path");
const childProcess = require("child_process");

// Lookup table for all platforms and binary distribution packages
const BINARY_DISTRIBUTION_PACKAGES = {
  darwin_x64: "mcpmock-darwin-x64",
  darwin_arm64: "mcpmock-darwin-arm64",
  linux_x64: "mcpmock-linux-x64",
  linux_arm64: "mcpmock-linux-arm64",
  freebsd_x64: "mcpmock-linux-x64",
  freebsd_arm64: "mcpmock-linux-arm64",
  win32_x64: "mcpmock-win32-x64",
  win32_arm64: "mcpmock-win32-arm64",
};

// Windows binaries end with .exe so we need to special case them.
const binaryName = process.platform === "win32" ? "mcpmock.exe" : "mcpmock";

// Determine package name for this platform
const platformSpecificPackageName =
  BINARY_DISTRIBUTION_PACKAGES[process.platform+"_"+process.arch];

function getBinaryPath() {
  try {
    // Resolving will fail if the optionalDependency was not installed
    return require.resolve(`@strowk/${platformSpecificPackageName}/bin/${binaryName}`);
  } catch (e) {
    return path.join(__dirname, "..", binaryName);
  }
}

module.exports.runBinary = function (...args) {
  childProcess.execFileSync(getBinaryPath(), args, {
    stdio: "inherit",
  });
};
