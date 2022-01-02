from typing import Iterable
from traits.MoveMax import MoveMax
from traits.Assess import Assess
from traits.Scope import Scope
from .Symbol import Symbol

def scope_processor(tokens: Iterable[str]):
    # TODO: Implement scoping interpreter
    token = ''
    while (token != Symbol("NEWLINE")):
        token = next(tokens)
    return tokens, Scope(lambda _: True)

def move_processor(tokens: Iterable[str]):
    origin = next(tokens)
    target = next(tokens)
    value = next(tokens)

    trait = MoveMax(origin, target, value)
    return tokens, trait

def assess_processor(tokens: Iterable[str]):
    field = next(tokens)
    return tokens, Assess(field)