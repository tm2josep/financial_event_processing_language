"""
Goal:

Create programming lanauge for processing financial/business events.
"""
from typing import Iterable
import pandas as pd
from Sequence import Sequence
from data_conatainers import AssessEvent, FinEvent
from interpreter.interpreter import manager
from traits.MoveMax import MoveMax
                
def dataframe_to_finevent_stream(df: pd.DataFrame):
    for _, row in df.iterrows():
        yield FinEvent(data=dict(row), scope_flag=True)

def main():
    event_dataframe = pd.read_csv('./test_data/standard_losses.csv')
    event_dataframe['liable'] = 0
    print(event_dataframe.dtypes)
    sequence = Sequence()
    traits = manager('./test_cases/scoped_move.fpl')
    for trait in traits:
        sequence.add_trait(trait)
    rows = sequence.process(dataframe_to_finevent_stream(event_dataframe))

    for row in rows:
        if (isinstance(row, AssessEvent)):
            print(row)
    pass

if (__name__ == "__main__"):
    main()