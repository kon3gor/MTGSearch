from dataclasses import dataclass
import json
from typing import List


def decapitalize(s: str) -> str:
    return s[0].lower() + s[1:]


@dataclass
class QueryParameter:
    url: str
    code: str
    expr: str
    type: str

    @staticmethod
    def from_dict(d: dict) -> 'QueryParameter':
        url = d["url"]
        code = d.get("code", url)
        expr = d.get("expr", code)
        type = d.get("type", "parameter")

        return QueryParameter(url, code, expr, type)


def create_token_const(qp: QueryParameter) -> str:
    return f"\t{qp.code} = \"{qp.expr}\""


def create_switch_branch(qp: QueryParameter) -> str:
    if qp.type == "parameter":
        return f"\tcase {qp.code}:\n\t\tq.{qp.code.capitalize()} = value{{v}}"
    else:
        return f"\tcase {qp.code}:\n\t\tq.{qp.code.capitalize()} = v"


def create_struct_entry(qp: QueryParameter) -> str:
    return f"\t{qp.code.capitalize()} {qp.type} `url:\"{qp.url},omitempty\"`"


def create_default_entry(qp: QueryParameter) -> str:
    if qp.type != "parameter":
        return None
    else:
        return f"\t{qp.code.capitalize()}: dp,"


def create_struct(fields: List[str]) -> str:
    f = "\n".join(fields)
    return f"type cardsQuery struct {{\n{f}\n}}"


def create_builder_func(branches: List[str]) -> str:
    f = "\n".join(branches)
    return f"func simpleCardsQuery(k, v string) cardsQuery {{\n\tq := dq\n\tswitch k {{\n{f}\n\t}}\n\treturn q\n}}"


def create_consts(values: List[str]) -> str:
    f = "\n".join(values)
    return f"const (\n{f}\n)"


def create_union_entry(qp: QueryParameter, union_func: str) -> str:
    if qp.type == "parameter":
        return f"\tif v := q.{qp.code.capitalize()}.{union_func}(o.{qp.code.capitalize()}); v != emptyParameter {{\n\t\tq.{qp.code.capitalize()} = v\n\t}}\n"
    else:
        return None


def create_union_func(entries: List[str], union_func: str) -> str:
    f = "\n".join(list(filter(lambda x: x, entries)))
    return f"func (q cardsQuery) {union_func}(o cardsQuery) cardsQuery {{\n{f}\n\treturn q\n}}"


with open("schema.json", "r") as file:
    data = json.load(file)
    fields = []
    branches = []
    consts = []
    default = []
    combine = []
    pipe = []
    for entry in data:
        qp = QueryParameter.from_dict(entry)
        consts.append(create_token_const(qp))
        branches.append(create_switch_branch(qp))
        fields.append(create_struct_entry(qp))
        default.append(create_default_entry(qp))
        combine.append(create_union_entry(qp, "combine"))
        pipe.append(create_union_entry(qp, "pipe"))

    struct = create_struct(fields)
    builder_func = create_builder_func(branches)
    consts = create_consts(consts)
    default_val = "\n".join(list(filter(lambda x: x, default)))
    combine_func = create_union_func(combine, "combine")
    pipe_func = create_union_func(pipe, "pipe")

    final = f"""package mtgsearch

{consts}

{struct}

var dp parameter = empty{{}}

var dq cardsQuery = cardsQuery{{
{default_val}
}}

{builder_func}

{combine_func}

{pipe_func}
    """

    with open("query.g.go", "w") as f:
        f.write(final)
