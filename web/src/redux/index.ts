import { LoginUserInfo } from "@/api/admin/sys/sys-user";
import { applyMiddleware, combineReducers, compose, legacy_createStore as createStore, Store } from "redux";
import { persistReducer, persistStore } from "redux-persist";
import storage from "redux-persist/lib/storage";
import reduxPromise from "redux-promise";
import reduxThunk from "redux-thunk";
import auth from "./modules/auth/reducer";
import breadcrumb from "./modules/breadcrumb/reducer";
import global from "./modules/global/reducer";
import menu from "./modules/menu/reducer";
import tabs from "./modules/tabs/reducer";

// 创建reducer(拆分reducer)
const reducer = combineReducers({
	global,
	menu,
	tabs,
	auth,
	breadcrumb
});

// redux 持久化配置
const persistConfig = {
	key: "redux-state",
	storage: storage
};
const persistReducerConfig = persistReducer(persistConfig, reducer);

// 开启 redux-devtools
const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;

// 使用 redux 中间件
const middleWares = applyMiddleware(reduxThunk, reduxPromise);

// 创建 store
const store: Store = createStore(persistReducerConfig, composeEnhancers(middleWares));

// 创建持久化 store
const persistor = persistStore(store);

export { persistor, store };
export type AppDispatch = typeof store.dispatch;
export const useDispatch = () => {
	return (_: any) => {};
};
export const useSelector = <T>(selector: (state: any) => T): T => {
	const uInfo: LoginUserInfo = store.getState().global.userInfo;
	return uInfo as T; // 返回一个空对象或指定类型的默认值
};

export type RootState = ReturnType<typeof store.getState>;
