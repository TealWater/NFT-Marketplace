import { writable } from "svelte/store";

//for loading n NFTs
export const collection_count = writable(50);
// @ts-ignore
export const wallet_state = writable( import("@web3-onboard/core").EIP1193Provider);