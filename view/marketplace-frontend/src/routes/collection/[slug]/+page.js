/* @type {import('./$types').PageLoad} */
import { PUBLIC_TRUSTED_URL } from '$env/static/public';
export async function load({fetch, params}) {

    // @ts-ignore
    const fetchCollectionNFTs = async (collection) => {
        const res = await fetch(`${PUBLIC_TRUSTED_URL}/getNFT?collection=${collection}`);
        const data = await res.json();
        return data.nfts;
    }
    return {
        opensea: fetchCollectionNFTs(params.slug) 
    };
};