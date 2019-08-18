import React from 'react';
import Header, { HeaderProps } from '../../Header';
import ShopDisplay, { ShopDisplayProps } from '../../ShopDisplay';
import CenterButton, { CenterButtonProps } from '../../CenterButton';

interface ShopSearchResultPageProps {
  header: HeaderProps;
  shops: ShopDisplayProps[];
  addButton: CenterButtonProps;
}

export default function ShopSearchResultPage(props: ShopSearchResultPageProps) {
  return (
    <form>
      <Header {...props.header} />
      {renderShops(props.shops)}
      <CenterButton
        {...props.addButton}
      />
    </form>
  );
}

const renderShops = (shops: ShopDisplayProps[]) => {
  return shops.map((shop, index) => (
    <ShopDisplay key={index} name={shop.name} />
  ));
};
