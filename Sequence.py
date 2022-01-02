from typing import Iterable
from traits.Trait import Trait
from data_conatainers import Event

class Sequence:
    def __init__(self):
        self.trait_list: Iterable[Trait] = []

    def add_trait(self, trait: Trait):
        self.trait_list.append(trait)

    def process(self, event_stack: Iterable[Event]):
        for trait in self.trait_list:
            event_stack = trait.process(event_stack)
        return event_stack