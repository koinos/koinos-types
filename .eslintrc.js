module.exports = {
  root: true,
  extends: [
    "airbnb-typescript/base",
    "plugin:@typescript-eslint/recommended",
    "plugin:@typescript-eslint/recommended-requiring-type-checking",
    "plugin:jest/recommended",
    "plugin:jest/style",
    "plugin:prettier/recommended",
  ],
  parserOptions: {
    project: "./tsconfig.eslint.json",
    "ecmaVersion": 2018,
  },
  rules: {
    "class-methods-use-this": "off",
    "import/no-cycle": "off",
    "prefer-template": "off",
    "no-bitwise": "off",
  },
};
