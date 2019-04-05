function add(state, block) {
  let properties =
    typeof block.properties !== "undefined" ? block.properties : null;

  state.blocks.push({
    index: state.blocks.length,
    type: block.type,
    name: block.name,
    properties: properties,
    value: null
  });
}

function change(state, { index, value }) {
  state.blocks[index].value = value;
}

function dragAndDrop(state, payload) {
  state.blocks = payload;
}

let setBlocks = (state, payload) => {
  state.blocks = payload.data
}

let setViewport = (state, payload) => {
  state.viewport = payload.data
}

export default {
  add,
  change,
  dragAndDrop,
  setBlocks,
  setViewport
};
