/* @type {import('./$types').PageLoad} */
import { PUBLIC_TRUSTED_URL } from '$env/static/public';
export async function load({ fetch, params }) {
    let contract = params.slug;
    let identifier = params.id;
    let chain = "ethereum"

    // @ts-ignore
    const fetchNFT = async (contract, identifier) => {
        const res = await fetch(`${PUBLIC_TRUSTED_URL}/getSingleNFT?address=${contract}&identifier=${identifier}`);
        const data = await res.json();
        return data.nft;
    }

    // @ts-ignore
    const fetchNFTEvents = async (contract, identifier) => {
        const res = await fetch(`${PUBLIC_TRUSTED_URL}/getSingleNFTEvents?address=${contract}&chain=${chain}&identifier=${identifier}&event_type=sale`);
        const data = await res.json();
        return data.asset_events;
    }

    return {
        nft: fetchNFT(contract, identifier),
        events: fetchNFTEvents(contract, identifier)
    };
};