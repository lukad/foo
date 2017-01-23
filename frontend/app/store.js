import {createStore, combineReducers, applyMiddleware} from "redux";
import thunk from "redux-thunk";
import promise from "redux-promise-middleware";

import math from "./reducers/mathReducer.js";
import user from "./reducers/userReducer.js";

const middlewares = [thunk, promise()];

if (process.env.NODE_ENV === "development") {
  const createLogger = require("redux-logger");
  const logger = createLogger();
  middlewares.push(logger);
}

export default createStore(
  combineReducers({math, user}),
  {},
  applyMiddleware(...middlewares)
);
