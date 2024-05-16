/* @type {import('./$types').PageServerLoad} */
export async function load({fetch}) {
    const fetchOpenSeaData = async () => {

        const res = await fetch(`http://localhost:8080/getEvents`);
        const data = await res.json();
        // console.log(data.asset_events);
        return data.asset_events;
    };

    return {
        opensea: fetchOpenSeaData()
    }
};