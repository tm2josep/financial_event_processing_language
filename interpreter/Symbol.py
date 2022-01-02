from dataclasses import dataclass
@dataclass
class Symbol:
    name: str
    def __repr__(self):
        return f"Symbol({self.name})"