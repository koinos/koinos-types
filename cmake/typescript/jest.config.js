module.exports = {
  preset: "ts-jest",
  testTimeout: 5000,
  maxConcurrency: 1,
  testEnvironment: "node",
  rootDir: ".",
  roots: ["<rootDir>/programs/"],
  testMatch: ["**/?(*.|*-)+(spec|test).ts"],
  transform: {
    "^.+\\.(t|j)s$": "ts-jest",
  },
  moduleFileExtensions: ["js", "json", "ts"],
};
