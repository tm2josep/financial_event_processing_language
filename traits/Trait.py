from typing import Iterable
from data_conatainers import FinEvent

class Trait:
    def __init__(self, name):
        self.name = name
    
    def process(self, event_stack: Iterable[FinEvent]) -> Iterable[FinEvent]:
        return event_stack