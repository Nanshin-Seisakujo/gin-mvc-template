module.exports = {
    root: true,
    extends: [
        "eslint:recommended",
        "plugin:import/errors",
        "plugin:import/warnings",
        "plugin:import/typescript",
        "plugin:@typescript-eslint/recommended",
        "plugin:@typescript-eslint/eslint-recommended",
        "plugin:prettier/recommended"
    ],
    plugins: ["@typescript-eslint", "import"],
    parser: "@typescript-eslint/parser",
    env: {
        browser: true,
        node: true,
        es6: true
    },
    settings: {
        react: {
            version: "detect"
        }
    },
    parserOptions: {
        sourceType: "module"
    },
    rules: {
        "@typescript-eslint/explicit-module-boundary-types": "off",
        "@typescript-eslint/no-empty-function": "off",
        "@typescript-eslint/ban-types": "off",
        "import/first": "warn",
        "import/no-duplicates": "warn",
        "import/order": "warn",
        "import/no-named-as-default": "off",
        "import/no-unresolved": "off"
    },
    overrides: [
        {
            files: ["*.js"],
            rules: { "@typescript-eslint/no-var-requires": ["off"] }
        }
    ]
};
