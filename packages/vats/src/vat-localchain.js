// @ts-check
import { Far } from '@endo/far';
import { makeDurableZone } from '@agoric/zone/durable.js';

import { prepareLocalChainTools } from './localchain.js';

export const buildRootObject = (_vatPowers, _args, baggage) => {
  const zone = makeDurableZone(baggage);
  const { makeLocalChain } = prepareLocalChainTools(zone.subZone('localchain'));

  return Far('LocalChainVat', {
    /**
     * Create a local chain that allows permissionlessly making fresh local
     * chain accounts, then using them to send chain queries and transactions.
     *
     * @param {import('./localchain.js').LocalChainPowers} powers
     */
    makeLocalChain(powers) {
      return makeLocalChain(powers);
    },
  });
};

/** @typedef {ReturnType<typeof buildRootObject>} LocalChainVat */
