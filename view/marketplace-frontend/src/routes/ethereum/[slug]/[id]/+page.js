/* @type {import('./$types').PageLoad} */
import { PUBLIC_TRUSTED_URL } from '$env/static/public';
export async function load({ fetch, params }) {
    let contract = params.slug;
    let identifier = params.id;

    // @ts-ignore
    const fetchNFT = async (contract, identifier) => {
        const res = await fetch(`${PUBLIC_TRUSTED_URL}/getSingleNFT?address=${contract}&identifier=${identifier}`);
        const data = await res.json();
        return data.nft;
    }
    return {
        nft: fetchNFT(contract, identifier)
    };
};