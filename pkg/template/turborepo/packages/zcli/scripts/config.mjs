import fs from "fs-extra";

const CONFIG_RESERVED_NAME = "zaf.config.json";

export const parseConfig = (path) => {
  try {
    let file = fs.readJsonSync(path, "utf-8");
    return file;
  } catch (e) {
    console.log("error in parse config", e);
    return null;
  }
};

// ignore folders where should be no config file
// add your own folders, if need be
const IGNORE_PACKAGES = ["zendesk"];

/**
 * Find config files - "zaf.config.json" in all packages
 * do not throw error if file is not found
 */
export const getAllConfigs = (packagesPath) => {
  const appsConfigsPaths = `${packagesPath}/apps-configs`;
  const packages = fs.readdirSync(appsConfigsPaths);
  const filtered = packages.filter((p) => !IGNORE_PACKAGES.includes(p));
  const configs = filtered.map((pkg) => {
    const cfgPath = `${appsConfigsPaths}/${pkg}/${CONFIG_RESERVED_NAME}`;
    return parseConfig(cfgPath);
  });
  return configs.filter((c) => c);
};
