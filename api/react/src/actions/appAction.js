import * as types from './actionTypes';
import axios from 'axios';

export function showResponse(response) {
    return {
        type: types.GET_RESPONSE,
        meaning: response
    };
}

export function getResponse(file, language) {
    return (dispatch, getState) => {
        var formData = new FormData();
        formData.append("file", file);
        axios.post('http://localhost:8080/?language=' + language, formData, {
            headers: {
            'Content-Type': 'multipart/form-data'
            }
        })
        .then(function (response) {
            dispatch(showResponse(response.data));
        })
        .catch(function (error) {
            dispatch(showResponse("Error"));
        });
    }
}