import * as types from '../actions/actionTypes';

const initialState = { }

export default function app(state = initialState, action) {
    switch (action.type) {
        case types.GET_RESPONSE:
            return {
                ...state,
                meaning: action.meaning
            };
        default:
            return state;
    }
}