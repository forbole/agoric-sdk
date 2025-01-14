/* eslint-disable */
/**
 * @file typesdef for the CoreEval environment
 *
 *   To use add this to the top of the proposal:
 *
 *   /// <reference types="@agoric/vats/src/core/core-eval-env" />
 *
 *   That directive has to be before imports, but this one's only useful in
 *   modules that have no imports or named exports.
 */

import type { VatData } from '@agoric/swingset-liveslots/src/vatDataTypes.js';
import type { E, Far, getInterfaceOf, passStyleOf } from '@endo/far';
import type { Assert, VirtualConsole } from 'ses';
import type { BootstrapModules } from './boot-chain.js';

// Provided by 'CORE_EVAL' handler in chain-behaviors.js
declare global {
  // bootstrap modules
  var behaviors: BootstrapModules['behaviors'];
  var utils: BootstrapModules['utils'];

  // @endo/far exports
  var E: E;
  var Far: Far;
  var getInterfaceOfFar: getInterfaceOfFar;
  var passStyleOfFar: passStyleOfFar;

  // endowments
  var VatData: VatData;
  var assert: Assert;

  // console is a VirtualConsole but this directive fails to override the extant global `console`
  // var console: VirtualConsole;

  // Base64 and URL are not available in all environments
}
