import * as React from 'react';
import { Switch } from 'react-router';
import { Link, Route } from 'react-router-dom';
import { Component } from 'react';
import TopPageContainer from './components/container/TopPageContainer';
import CreateHousePageContainer from './components/container/CreateHousePageContainer';
import NotFound from './components/presentation/NotFound';
import { Box } from '@material-ui/core';
import AdminHouseTopPageContainer from './components/container/admin/HouseTopPageContainer';
import MemberHouseTopPageContainer from './components/container/member/HouseTopPageContainer';
import EventCreatePageContainer from './components/container/admin/EventCreatePageContainer';
import EventDetailPageContainer from './components/container/admin/EventDetailPageContainer';
import QuestionnaireResultPageContainer from './components/container/admin/QuestionnaireResultPageContainer';
import QuestionnaireCreatePageContainer from './components/container/admin/QuestionnaireCreatePageContainer';
import QuestionnaireDetailPageContainer from './components/container/admin/QuestionnaireDetailPageContainer';
import ShopSearchSettingPageContainer from './components/container/admin/ShopSearchSettingPageContainer';
import ShopSearchResultPageContainer from './components/container/admin/ShopSearchResultPageContainer';
import RegisterPageContainer from './components/container/RegisterPageContainer';
import ReceiveQuestionnairesPageContainer from './components/container/member/ReceiveQuestionnairesPageContainer';
import DateDescriptionFirstPageContainer from './components/container/member/DateDescriptionFirstPageContainer';

// ルート名一覧
export const ROUTE_ROOT = '/';
export const ROUTE_USER_REGISTER = '/users/register';
export const ROUTE_HOUSE_CREATE = '/houses/create';
export const ROUTE_HOUSE_LOGIN = '/houses/login';
export const ROUTE_ADMIN_HOUSE_TOP = '/admin/houses/top'; // 仮
export const ROUTE_ADMIN_HOUSE_EVENT_CREATE = '/admin/houses/1/events/create';
export const ROUTE_ADMIN_HOUSE_EVENT_DETAIL = '/admin/houses/1/events/1';
export const ROUTE_ADMIN_HOUSE_EVENT_QUES =
  '/admin/houses/1/events/1/questionnaires';
export const ROUTE_ADMIN_HOUSE_EVENT_QUES_CREATE =
  '/admin/houses/1/events/1/questionnaires/create';
export const ROUTE_ADMIN_HOUSE_EVENT_QUES_DETAIL =
  '/admin/houses/1/events/1/questionnaires/1';

export const ROUTE_ADMIN_HOUSE_EVENT_QUES_SHOP_SEARCH_SETTING =
  '/admin/houses/1/events/1/questionnaires/1/shop/search';
export const ROUTE_ADMIN_HOUSE_EVENT_QUES_SHOP_SEARCH_RESULT =
  '/admin/houses/1/events/1/questionnaires/1/shop/search/result';

export const ROUTE_MEMBER_HOUSE_TOP = '/member/house/top';
export const ROUTE_MEMBER_QUESTIONNAIRES_RECEIVED =
  '/member/house/1/questionnaires/received';
export const ROUTE_MEMBER_DATE_DESC_FIRST =
  '/member/house/1/questionnaires/1/descriptions';

export class Routes extends Component<{}, {}> {
  render() {
    return (
      <Box m={0}>
        <h1>デバッグ用ルーティング</h1>
        <li>
          <Link to={ROUTE_ROOT}>Home</Link>
        </li>
        <li>
          <Link to={ROUTE_USER_REGISTER}>ユーザー新規作成</Link>
        </li>
        <li>
          <Link to={ROUTE_HOUSE_CREATE}>ハウス作成画面</Link>
        </li>

        <li>
          <Link to={ROUTE_ADMIN_HOUSE_TOP}>管理者ハウストップページ</Link>
        </li>
        <li>
          <Link to={ROUTE_ADMIN_HOUSE_EVENT_CREATE}>
            管理者イベント作成ページ
          </Link>
        </li>
        <li>
          <Link to={ROUTE_ADMIN_HOUSE_EVENT_DETAIL}>管理者イベント詳細</Link>
        </li>
        <li>
          <Link to={ROUTE_ADMIN_HOUSE_EVENT_QUES}>アンケート一覧（結果）</Link>
        </li>
        <li>
          <Link to={ROUTE_ADMIN_HOUSE_EVENT_QUES_CREATE}>アンケート作成</Link>
        </li>
        <li>
          <Link to={ROUTE_ADMIN_HOUSE_EVENT_QUES_DETAIL}>アンケート詳細</Link>
        </li>
        <li>
          <Link to={ROUTE_ADMIN_HOUSE_EVENT_QUES_SHOP_SEARCH_SETTING}>
            アンケート用店検索設定
          </Link>
        </li>
        <li>
          <Link to={ROUTE_ADMIN_HOUSE_EVENT_QUES_SHOP_SEARCH_RESULT}>
            店検索結果
          </Link>
        </li>
        <li>
          <Link to={ROUTE_MEMBER_HOUSE_TOP}>メンバーハウストップ</Link>
        </li>
        <li>
          <Link to={ROUTE_MEMBER_QUESTIONNAIRES_RECEIVED}>
            受信アンケート一覧
          </Link>
        </li>
        <li>
          <Link to={ROUTE_MEMBER_DATE_DESC_FIRST}>日付アンケート説明1</Link>
        </li>
        <Switch>
          <Route exact path={ROUTE_ROOT} component={TopPageContainer} />
          <Route
            exact
            path={ROUTE_USER_REGISTER}
            component={RegisterPageContainer}
          />
          <Route
            exact
            path={ROUTE_HOUSE_CREATE}
            component={CreateHousePageContainer}
          />
          <Route
            exact
            path={ROUTE_ADMIN_HOUSE_TOP}
            component={AdminHouseTopPageContainer}
          />
          <Route
            exact
            path={ROUTE_ADMIN_HOUSE_EVENT_CREATE}
            component={EventCreatePageContainer}
          />
          <Route
            exact
            path={ROUTE_ADMIN_HOUSE_EVENT_DETAIL}
            component={EventDetailPageContainer}
          />
          <Route
            exact
            path={ROUTE_ADMIN_HOUSE_EVENT_QUES}
            component={QuestionnaireResultPageContainer}
          />
          <Route
            exact
            path={ROUTE_ADMIN_HOUSE_EVENT_QUES_CREATE}
            component={QuestionnaireCreatePageContainer}
          />
          <Route
            exact
            path={ROUTE_ADMIN_HOUSE_EVENT_QUES_DETAIL}
            component={QuestionnaireDetailPageContainer}
          />
          <Route
            exact
            path={ROUTE_ADMIN_HOUSE_EVENT_QUES_SHOP_SEARCH_SETTING}
            component={ShopSearchSettingPageContainer}
          />
          <Route
            exact
            path={ROUTE_ADMIN_HOUSE_EVENT_QUES_SHOP_SEARCH_RESULT}
            component={ShopSearchResultPageContainer}
          />
          <Route
            exact
            path={ROUTE_MEMBER_HOUSE_TOP}
            component={MemberHouseTopPageContainer}
          />
          <Route
            exact
            path={ROUTE_MEMBER_QUESTIONNAIRES_RECEIVED}
            component={ReceiveQuestionnairesPageContainer}
          />
          <Route
            exact
            path={ROUTE_MEMBER_DATE_DESC_FIRST}
            component={DateDescriptionFirstPageContainer}
          />
          <Route component={NotFound} />
        </Switch>
      </Box>
    );
  }
}
