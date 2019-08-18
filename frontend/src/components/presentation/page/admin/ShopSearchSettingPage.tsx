import React from 'react';
import SimpleSelect, { SimpleSelectProps } from '../../SimpleSelect';
import Header, { HeaderProps } from '../../Header';
import TextInputField, { TextInputFieldProps } from '../../TextInputField';
import CenterButton, { CenterButtonProps } from '../../CenterButton';

interface ShopSearchSettingPageProps {
  header: HeaderProps;
  regionSelect: SimpleSelectProps;
  genreSelect: SimpleSelectProps;
  ratingSelect: SimpleSelectProps;
  budgetInput: TextInputFieldProps;
  searchButton: CenterButtonProps;
}

export default function ShopSearchSettingPage(
  props: ShopSearchSettingPageProps
) {
  return (
    <form>
      <Header {...props.header} />
      <SimpleSelect {...props.regionSelect} />
      <SimpleSelect {...props.genreSelect} />
      <SimpleSelect {...props.ratingSelect} />
      <TextInputField {...props.budgetInput} />
      <CenterButton {...props.searchButton} />
    </form>
  );
}
