/* @type {import('./$types').PageLoad} */
import { PUBLIC_TRUSTED_URL } from '$env/static/public';
import { collection_count } from '$lib/stores/store.js';
export async function load({fetch}) {

    let amnt = 0;
    //Will allow amnt to be updated whenever the store is updated.
    collection_count.subscribe((prev_val) => amnt = prev_val );
    let str = JSON.stringify(amnt);

    // @ts-ignore
    const fetchTopOpenseaCollections = async (str) => {
        const res = await fetch(`${PUBLIC_TRUSTED_URL}/getTopCollections?limit=${str}`);
        const data = await res.json();
        return data.collections;
    }

    return {
        opensea : fetchTopOpenseaCollections(amnt)
    };
};