/** @file Ambient exports until https://github.com/Agoric/agoric-sdk/issues/6512 */
/** @see {@link /docs/typescript.md} */
/* eslint-disable -- doesn't understand .d.ts */

// XXX also explicit exports because `@agoric/ertp` top level confuses the type and value of `AssetKind`
export * from './src/types.js';

import {
  Amount as _Amount,
  Brand as _Brand,
  Issuer as _Issuer,
  IssuerKit as _IssuerKit,
  Mint as _Mint,
  AssetKind as _AssetKind,
  SetValue as _SetValue,
  NatValue as _NatValue,
  DisplayInfo as _DisplayInfo,
  AdditionalDisplayInfo as _AdditionalDisplayInfo,
  Payment as _Payment,
  Purse as _Purse,
} from './src/types.js';
declare global {
  export {
    _Amount as Amount,
    _Brand as Brand,
    _Issuer as Issuer,
    _IssuerKit as IssuerKit,
    _Mint as Mint,
    _AssetKind as AssetKind,
    _SetValue as SetValue,
    _NatValue as NatValue,
    _DisplayInfo as DisplayInfo,
    _AdditionalDisplayInfo as AdditionalDisplayInfo,
    _Payment as Payment,
    _Purse as Purse,
  };
}
