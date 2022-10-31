import createRouter from "../framework/router";
import createStore from "../framework/store";
import routes from "./router";
import storeObj from "./store/index";

const store = createStore(storeObj);
const router = createRouter(routes);

export { store, router };