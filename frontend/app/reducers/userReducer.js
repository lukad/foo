export default (state = {name: "Jamal", age: 23}, action) => {
  switch (action.type) {
  case "SET_NAME":
  case "SET_NAME_FULFILLED":
    return {
      ...state,
      name: action.payload
    };
  case "SET_NAME_PENDING":
    return {
      ...state,
      name: "pending..."
    };
  case "SET_AGE":
    return {
      ...state,
      age: action.payload
    };
  default:
    return state;
  }
};
