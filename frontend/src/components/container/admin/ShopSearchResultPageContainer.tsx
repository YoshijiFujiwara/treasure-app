import React, {Component} from 'react';
import ShopSearchResultPage from "../../presentation/page/admin/ShopSearchResultPage";
import {ShopDisplayProps} from "../../presentation/ShopDisplay";
import {ROUTE_ADMIN_HOUSE_EVENT_QUES_DETAIL} from "../../../Routes";

export default class ShopSearchResultPageContainer extends Component {
    render() {
        return (
            <ShopSearchResultPage
                header={{title: '店検索結果'}}
                shops={dummyShops}
                addButton={{title: '店をアンケートに追加する', linkTo: ROUTE_ADMIN_HOUSE_EVENT_QUES_DETAIL}}
            />
        )
    }
}

const dummyShops: ShopDisplayProps[] = [
    {
        name: 'ホゲホゲラーメン'
    },
    {
        name: 'ふがふがそば'
    },
    {
        name: 'ホゲホゲマック'
    }
];
