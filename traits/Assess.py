from .Trait import Trait
from typing import Iterable
from data_conatainers import AssessEvent, FinEvent
from copy import copy

class Assess(Trait):
    def __init__(self, target: str):
        super().__init__('move')

        self.target = target
    
    def process(self, event_stack: Iterable[FinEvent]) -> Iterable[AssessEvent]:
        assessed_value = 0
        for event in event_stack:
            data = copy(event.data)
            if (self.target not in data):
                data[self.target] = 0
            assessed_value += data[self.target]
        yield AssessEvent(scope_flag=False, field=self.target, value=assessed_value)