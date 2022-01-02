from .Trait import Trait
from typing import Iterable
from data_conatainers import FinEvent


class Scope(Trait):
    def __init__(self, condition):
        super().__init__('move')

        self.condition = condition

    def process(self, event_stack: Iterable[FinEvent]) -> Iterable[FinEvent]:
        return map(
            lambda event: FinEvent(
                data=event.data,
                scope_flag=self.condition(event.data)
            ),
            event_stack
        )
