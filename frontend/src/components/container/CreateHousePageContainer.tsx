import React, {Component} from "react";
import CreateHousePage from '../presentation/page/CreateHousePage';
import {ROUTE_ADMIN_HOUSE_TOP} from "../../Routes";

export default class CreateHousePageContainer extends Component {
    render() {
        return (
            <div>
                <CreateHousePage headerTitle={{title: 'ハウスを新規作成する'}} centerButton={{title: '家族名', linkTo: ROUTE_ADMIN_HOUSE_TOP}} textLabel="次へ" />
            </div>
        )
    }
}
