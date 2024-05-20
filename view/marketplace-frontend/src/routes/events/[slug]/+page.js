/* @type {import('./$types').PageServerLoad} */
import { PUBLIC_TRUSTED_URL } from '$env/static/public';
export async function load({ fetch, params }) {
    // @ts-ignore
    const fetchOpenSeaData = async (collection) => {

        const res = await fetch(`${PUBLIC_TRUSTED_URL}/getEvents?collection=${collection}`);
        const data = await res.json();
        // console.log(data.asset_events);
        return data.asset_events;
    };

    return {
        opensea: fetchOpenSeaData(params.slug),
        collection: params.slug
    }
};