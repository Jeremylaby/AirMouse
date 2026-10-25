// https://docs.expo.dev/guides/using-eslint/
const { defineConfig } = require('eslint/config');
const expoConfig = require('eslint-config-expo/flat');
const eslintConfigPrettier = require('eslint-config-prettier/flat');
const reactNative = require('eslint-plugin-react-native');

module.exports = defineConfig([
  expoConfig,
  eslintConfigPrettier,
  {
    plugins: {
      'react-native': reactNative,
    },
    rules: {
      // React Native
      'react-native/no-unused-styles': 'error',
      'react-native/no-single-element-style-arrays': 'error',
      'react-native/no-inline-styles': 'warn',
      'react-native/no-color-literals': 'warn',

      // React Hooks
      'react-hooks/rules-of-hooks': 'error',

      // General
      'no-console': 'warn',
      'prefer-const': 'error',
      'no-var': 'error',
    },
  },
  {
    ignores: ['dist/*'],
  },
]);
