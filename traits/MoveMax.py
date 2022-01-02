from .Trait import Trait
from typing import Iterable
from data_conatainers import FinEvent
from copy import copy

class MoveMax(Trait):
    def __init__(self, origin: str, target: str, value: float):
        super().__init__('move')

        self.origin = origin
        self.target = target
        self.value = float(value)
    
    def process(self, event_stack: Iterable[FinEvent]) -> Iterable[FinEvent]:
        for event in event_stack:
            data, scope_flag = event.data, event.scope_flag
            print(event)

            if (scope_flag):
                new_data = copy(data)
                usable_value = min(float(new_data[self.origin]), self.value)
                new_data[self.origin] -= usable_value
                new_data[self.target] += usable_value

            yield FinEvent(data=new_data, scope_flag=scope_flag)