/* @type {import('./$types').PageServerLoad} */
export async function load({fetch, params}) {
    // @ts-ignore
    const fetchOpenSeaData = async (collection) => {

        const res = await fetch(`http://localhost:8080/getEvents?collection=${collection}`);
        const data = await res.json();
        // console.log(data.asset_events);
        return data.asset_events;
    };

    return {
        opensea: fetchOpenSeaData(params.slug)
    }
};