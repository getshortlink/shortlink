const path = require('path');
const copyPlugin = require('copy-webpack-plugin');

module.exports = {
    mode: 'production',
    entry: {
        background: path.resolve(__dirname, "..", "src", "worker","background.ts")
    },
    output: {
        path: path.join(__dirname, "../dist"),
        filename: "[name].js"
    },
    resolve: {
        extensions: [".ts", ".js"]
    },
    module: {
        rules: [
           {
              test: /\.tsx?$/,
              loader: "ts-loader",
              exclude: /node_modules/,
           },
        ],
     },
     plugins: [
        new copyPlugin({
           patterns: [{from: ".", to: ".", context: "public"}]
        }),
     ],
}