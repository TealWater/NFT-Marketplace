import { json } from "@sveltejs/kit"

/**
 * @param {string} message
 */
export function CreateNFTEvent(message){

    let nft_event = {
        collection: '',
        event: '',
        chain: '',
        timestamp: '',
        quantity: 0
    }

    const event = JSON.parse(message)
    nft_event.event = event.event;
      console.log("#0: ", event);
    //const event_payload = JSON.parse(event.payload);
    const event_payload = event.payload;
    console.log("#1: ", event_payload);
    const event_payload_data = event_payload;
    console.log("#2: ", event_payload_data);
    nft_event.chain = event_payload_data.chain;
    console.log("#333: ", event_payload_data.payload.collection.slug);
    nft_event.collection = event_payload_data.payload.collection.slug
    nft_event.timestamp = event_payload.payload.event_timestamp;
    nft_event.quantity = event_payload.payload.quantity;

    return nft_event;
}