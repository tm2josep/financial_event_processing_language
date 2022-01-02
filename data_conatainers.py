from dataclasses import dataclass

@dataclass
class Event:
    scope_flag: bool

@dataclass
class FinEvent(Event):
    data: dict

@dataclass
class AssessEvent(Event):
    field: str
    value: float