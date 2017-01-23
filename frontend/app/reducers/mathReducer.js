export default (state = {result: 1, lastValues: []}, action) => {
  switch (action.type) {
  case "ADD":
    return {
      ...state,
      result: state.result + action.payload,
      lastValues: [...state.lastValues, action.payload]
    };
  case "SUB":
    return {
      ...state,
      result: state.result - action.payload,
      lastValues: [...state.lastValues, action.payload]
    };
  default:
    return state;
  }
};
