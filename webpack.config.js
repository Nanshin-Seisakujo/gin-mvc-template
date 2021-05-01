const path = require("path");
const TsconfigPathsPlugin = require("tsconfig-paths-webpack-plugin");
const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const FixStyleOnlyEntriesPlugin = require("webpack-fix-style-only-entries");
const WebpackAssetsManifest = require("webpack-assets-manifest");

const env = process.env.NODE_ENV || "development";
const isProd = env === "production";

module.exports = {
    mode: env,
    cache: isProd ? false : true,
    entry: {
        main: path.resolve(__dirname, "src/typescript/main.ts"),
        css: path.resolve(__dirname, "src/sass/main.scss")
    },
    output: {
        path: path.resolve(__dirname, "static"),
        filename: isProd
            ? "js/[name]-[contenthash].bundle.js"
            : "js/[name].bundle.js"
    },
    devtool: isProd ? false : "inline-source-map",
    module: {
        rules: [
            {
                test: /\.js$/,
                exclude: /core-js/,
                use: [
                    {
                        loader: "babel-loader",
                        options: {
                            presets: ["@babel/preset-env"]
                        }
                    }
                ]
            },
            {
                test: /\.ts?$/,
                exclude: /node_modules/,
                use: [
                    {
                        loader: "ts-loader"
                    }
                ]
            },
            {
                test: /\.scss$/i,
                use: [
                    MiniCssExtractPlugin.loader,
                    "css-loader",
                    "postcss-loader",
                    "sass-loader"
                ]
            }
        ]
    },
    resolve: {
        extensions: [".ts", ".js", ".css", ".scss"],
        plugins: [
            new TsconfigPathsPlugin({
                configFile: "./tsconfig.json"
            })
        ],
        modules: [path.resolve(__dirname, "src/typescript"), "node_modules"]
    },
    plugins: [
        new MiniCssExtractPlugin({
            filename: isProd ? "css/main-[contenthash].css" : "css/main.css"
        }),
        new FixStyleOnlyEntriesPlugin(),
        new WebpackAssetsManifest({
            output: "../.data/manifest.json",
            publicPath: true,
            writeToDisk: true,
            entrypoints: true
        })
    ]
};
