# https://rosettacode.org/wiki/Burrows%E2%80%93Wheeler_transform#Python
"""Apply Burrows-Wheeler transform to input string."""
def bwt(s):
    assert "\002" not in s and "\003" not in s, "Input string cannot contain STX and ETX characters"
    s = "\002" + s + "\003"  # Add start and end of text marker
    table = sorted(s[i:] + s[:i] for i in range(len(s)))  # Table of rotations of string
    last_column = [row[-1:] for row in table]  # Last characters of each row
    return "".join(last_column)  # Convert list of characters into string

"""Apply inverse Burrows-Wheeler transform."""
def ibwt(r):
    table = [""] * len(r)  # Make empty table
    for i in range(len(r)):
        table = sorted(r[i] + table[i] for i in range(len(r)))  # Add a column of r
    s = [row for row in table if row.endswith("\003")][0]  # Find the correct row (ending in ETX)
    return s.rstrip("\003").strip("\002")  # Get rid of start and end markers

def compress(s):
    t = bwt(s)

    return t

if __name__ == "__main__":
    sequence = "AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC"

    compressed = compress(sequence)

    print(compressed)
