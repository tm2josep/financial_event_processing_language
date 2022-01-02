from types import FunctionType
from typing import Union
import os
from interpreter.Symbol import Symbol
from .processors import move_processor, assess_processor, scope_processor

NEWLINE = Symbol('NEWLINE')
ITERATOR_COMPLETE = Symbol('ITERATOR_COMPLETE')

def processor_selector(token: str) -> FunctionType:
    if (token == 'move'):
        return move_processor
    if (token == 'assess'):
        return assess_processor
    if (token == 'scope'):
        return scope_processor

def manager(source_file_url: Union[str, bytes, os.PathLike]):
    tokens = read_tokens_from_file(source_file_url)
    traits = []
    while (True):
        token = next(tokens, ITERATOR_COMPLETE)
        if (token == ITERATOR_COMPLETE):
            break
        if (token == NEWLINE):
            continue

        processor = processor_selector(token)
        tokens, trait = processor(tokens)
        traits.append(trait)

    return traits

def read_tokens_from_file(source_file_url: Union[str, bytes, os.PathLike]):
    with open(source_file_url) as f:
        while (line := f.readline()):
            tokens = line.split(' ')
            tokens = map(lambda t: t.strip(), tokens)
            tokens = filter(lambda t: len(t) > 0, tokens)
            yield from tokens
            yield NEWLINE


if (__name__ == "__main__"):
    traits = manager(os.path.join('.', 'test_cases', 'simple_move.fpl'))

