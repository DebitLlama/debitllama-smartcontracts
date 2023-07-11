// rollup.config.js
import { nodeResolve } from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import json from '@rollup/plugin-json';
import terser from '@rollup/plugin-terser';
import nodePolyfills from 'rollup-plugin-polyfill-node';
import replace from "@rollup/plugin-replace";

export default {
    input: 'lib/directDebit.js',
    output: {
        file: 'directdebit_bundle.js',
        format: 'iife',
        name: "directdebitlib",
        globals: {
            process: { browser: true }
        }
    },
    plugins: [nodeResolve({ browser: true }), commonjs(), json(), terser(), nodePolyfills(), replace({
        preventAssignment: true,
        "process.browser": true
    })]
};