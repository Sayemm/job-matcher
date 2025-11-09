from dataclasses import dataclass

@dataclass
class Resume:
    text: str

    def is_valid(self) -> bool:
        if not self.text:
            return False
        if len(self.text.strip()) < 50:
            return False
        return True
    
    def get_text_for_matching(self) -> str:
        return self.text.strip()