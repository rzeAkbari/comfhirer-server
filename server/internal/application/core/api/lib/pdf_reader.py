import io
import pickle
import PyPDF2
import re
import json

path = "/Users/raziehakbari/raz-project/DNF/comfhirer-server/server/internal/application/core/api/lib/data"

try:
    with open(path+'/principio_attivo_pattern.pkl', 'rb') as f:
        principio_attivo_pattern = pickle.load(f)
    with open(path+'/Farmaco_pattern.pkl', 'rb') as f:
        Farmaco_pattern = pickle.load(f)
    with open(path+'/AIC_pattern.pkl', 'rb') as f:
        AIC_pattern = pickle.load(f)
except EnvironmentError as err:
    print(err)


def Do(pdf:bytearray):
    extracted_text = extract_text(pdf)
    cleaned_text = clean_text(extracted_text)
    nome = find_text(cleaned_text, "NOME:")
    aic = find_pattern(AIC_pattern, cleaned_text)
    farmaco = find_pattern(Farmaco_pattern, cleaned_text)
    principio_attivo = find_pattern(principio_attivo_pattern, cleaned_text)
    confezione_di_riferimento = find_text (cleaned_text, principio_attivo)
    Dict = {'nome': nome,
        'aic_code': aic,
        'farmaco': farmaco,
        'principio_attivo': principio_attivo,
        'Confezione_di_riferimento': confezione_di_riferimento
        }
    return json.dumps(Dict)


def extract_text(pdf:bytearray):
    with io.BytesIO(pdf) as pdf:
        reader=PyPDF2.PdfReader(pdf, strict=False)
        pdf_text = []

        for page in reader.pages:
            content = page.extract_text()
            pdf_text.append(content)

        return pdf_text


def clean_text(txt):
    txt[0] = re.sub(r'[_].*', ' ', txt[0], flags=re.MULTILINE)
    text = txt[0].split('\n')
    text = [x for x in text if x != '  ']
    text = [x for x in text if x != ' ']
    return text


def find_text(text:list, word:str):
    spl_word = word
    matching = [s for s in text if spl_word in s]
    name=matching[0][matching[0].find(spl_word)+len(spl_word):]

    return name


def find_pattern (pattern, text:list):
    pattern_list = []
    for x in text:
        patternRegex = re.compile(str(pattern))
        try:
            pt = patternRegex.search(x)
            result = pt.group()
            pattern_list.append(result)
        except:
            pass
    return pattern_list[0]