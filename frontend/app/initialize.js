import ReactDOM from "react-dom";
import React from "react";
import {Provider} from "react-redux";
import store from "./store";

import App from "containers/App";

const load = () => {
  ReactDOM.render(
    <Provider store={store}>
      <App />
    </Provider>,
    document.querySelector("#app")
  );
};

if (document.readyState !== "complete") {
  document.addEventListener("DOMContentLoaded", load);
} else {
  load();
}
