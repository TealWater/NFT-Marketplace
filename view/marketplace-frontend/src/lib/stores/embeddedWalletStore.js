/*
this file makes use of the Web3-Onboard JS library 
(https://onboard.blocknative.com/)

For an object / value to be a valid store it must implment the `.subscribe` method
to satisfy the store contract.
`.subscribe` method must also return an unsubscribe func.
*/
import { PUBLIC_INFURA_ID } from '$env/static/public';
import Onboard from '@web3-onboard/core';
import injectedModule from '@web3-onboard/injected-wallets';
import coinbaseModule from '@web3-onboard/coinbase';
import phantomModule from '@web3-onboard/phantom';
import metamaskSDK from '@web3-onboard/metamask';
import { Buffer } from 'buffer';

// @ts-ignore
globalThis.Buffer = Buffer
const injected = injectedModule();
const coinbase = coinbaseModule();
const phantom = phantomModule();
const metamask = metamaskSDK({
    options: {
        extensionOnly: false,
        dappMetadata: {
            name: 'Demo Web3Onboard'
        }
    }
});

const wallets = [
    metamask,
    coinbase,
    phantom,
    injected
];

const chains = [
    {
        id: '0x1',
        token: 'ETH',
        label: 'Ethereum Mainnet',
        rpcUrl: `https://mainnet.infura.io/v3/${PUBLIC_INFURA_ID}`
    },
    {
        id: 11155111,
        token: 'ETH',
        label: 'Sepolia',
        rpcUrl: 'https://rpc.sepolia.org/'
    },
    {
        id: '0x13881',
        token: 'MATIC',
        label: 'Polygon - Mumbai',
        rpcUrl: 'https://matic-mumbai.chainstacklabs.com'
    },
    {
        id: '0x38',
        token: 'BNB',
        label: 'Binance',
        rpcUrl: 'https://bsc-dataseed.binance.org/'
    },
    {
        id: '0xA',
        token: 'OETH',
        label: 'OP Mainnet',
        rpcUrl: 'https://mainnet.optimism.io'
    },
    {
        id: '0xA4B1',
        token: 'ARB-ETH',
        label: 'Arbitrum',
        rpcUrl: 'https://rpc.ankr.com/arbitrum'
    },
    {
        id: '0xa4ec',
        token: 'ETH',
        label: 'Celo',
        rpcUrl: 'https://1rpc.io/celo'
    },
    {
        id: 666666666,
        token: 'DEGEN',
        label: 'Degen',
        rpcUrl: 'https://rpc.degen.tips'
    }
]

const appMetadata = {
    name: 'Connect Wallet',
    icon: '<svg>My App Icon</svg>',
    description: 'Example showcasing how to connect a wallet.',
    recommendedInjectedWallets: [
        { name: 'MetaMask', url: 'https://metamask.io' },
        { name: 'Coinbase', url: 'https://wallet.coinbase.com/' },
        { name: 'Phantom', url: 'https://phantom.app/' }
    ]
}

const onboard = Onboard({
    wallets,
    chains,
    appMetadata
});

export default onboard