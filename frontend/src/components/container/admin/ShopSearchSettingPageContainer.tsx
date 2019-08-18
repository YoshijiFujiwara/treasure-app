import React, {Component} from 'react';
import ShopSearchSettingPage from "../../presentation/page/admin/ShopSearchSettingPage";

export default class ShopSearchSettingPageContainer extends Component {
    render() {
        return (
            <ShopSearchSettingPage
                header={{title: '店検索設定'}}
                regionSelect={{label: '地域'}}
                genreSelect={{label: 'ジャンル'}}
                ratingSelect={{label: '評価'}}
                budgetInput={{label: '予算'}}
                searchButton={{title: 'この条件で検索', linkTo: 'hogehoge'}}
            />
        )
    }
}
