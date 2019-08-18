import React from 'react';
import Header, { HeaderProps } from '../../Header';
import QuestionnaireCard, {
  QuestionnaireCardProps
} from '../../QuestionnaireCard';

interface QuestionnaireResultPageProps {
  header: HeaderProps;
  questionnires: QuestionnaireCardProps[];
}

export default function QuestionnaireResultPage(
  props: QuestionnaireResultPageProps
) {
  return (
    <div>
      <Header {...props.header} />
      {renderQuestionniares(props.questionnires)}
    </div>
  );
}

const renderQuestionniares = (quesionniares: QuestionnaireCardProps[]) => {
  return quesionniares.map((quesionniare, index) => (
    <QuestionnaireCard key={index} name={quesionniare.name} />
  ));
};
