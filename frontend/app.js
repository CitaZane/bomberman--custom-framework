import createRouter from "../framework/router";
import createStore from "../framework/store"
import router from "./router/index";
import storeObj from "./store/index";

const store = createStore(storeObj);
createRouter(router);

export { store }

