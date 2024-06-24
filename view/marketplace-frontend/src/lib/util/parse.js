/**
 * @param {string} message
 */
export function CreateNFTEventFromSocket(message) {

  let nft_event = {
    collection: '',
    event: '',
    chain: '',
    timestamp: '',
    quantity: 0,
    maker: '',
    taker: ''
  }

  const event = JSON.parse(message)
  nft_event.event = event.event;
  // console.log("#0: ", event);
  //const event_payload = JSON.parse(event.payload);
  const event_payload = event.payload;
  // console.log("#1: ", event_payload);
  const event_payload_data = event_payload;
  // console.log("#2: ", event_payload_data);
  nft_event.chain = event_payload_data.chain;
  // console.log("#333: ", event_payload_data.payload.collection.slug);
  // nft_event.collection = event_payload_data.payload.collection.slug
  nft_event.collection = event_payload_data.payload.item.metadata.name;

  //convert from UTC to unix time
  let time = new Date(event_payload.payload.event_timestamp).getTime() / 1000;
  nft_event.timestamp = time.toFixed(0)
  // console.log("socket info: ", nft_event.timestamp);

  nft_event.quantity = event_payload.payload.quantity;
  nft_event.maker = event_payload.payload.maker.address;

  return nft_event;
}