

def getLines():
	s = open('status.txt').read()

	sections = []
	for line in s.splitlines():
		line = line.strip()
		if len(line) > 0:
			sections.append(line)
	sections.reverse()
	return sections

def parse():
	lines = getLines()
	
	parts = []
	try:
		while len(lines) > 1:
			line = lines.pop()
			if line.startswith("-----"):
				parts.append(parsePart(lines))
	except Exception as ex:
		print ex

	return parts

import re

def parsePart(lines):
	line = lines.pop()
	if '|' in line and not '*' in line:
		item = {}
		item['name'] = line[1:len(line)-1]
		parts = parseSection(lines)
		if parts is None or len(parts) != 2:
			print "line = ", line
			raise Exception("parse error")
		match = re.search(r'\<[^>]+', parts[1])
		if match:
			item['description'] = parts[1].replace(match.group(0) + ">", "https://mariadb.com" + match.group(0)[1:])
		else:
			item['description'] = parts[1]
		parts = parseSection(lines)
		if parts is None or len(parts) != 2:
			print "parts error: ", parts
			raise Exception("parse error")
		item['scope'] = parts[1]
		parts = parseSection(lines)
		if parts is None or len(parts) != 2:
			print "parts error2: ", parts
			raise Exception("parse error")
		item['type'] = parts[1]
		return item
		


def parseSection(lines):
	
	sectionFound = False
	while  len(lines) > 1:
		line = lines.pop()
		if line.startswith("----"):
			if sectionFound:
				return field, data
			continue
		if line.startswith("* *"):
			if sectionFound:
				lines.append(line)
				return field, data
			sectionFound = True
			parts = line[3:].split(":* ")
			if len(parts) == 1:
				return parts[0], ""
			field = parts[0]
			data = parts[1]
		elif sectionFound:
			data += " " + line


def to_camelcase(s):
	return re.sub(r'(?!^)_([a-zA-Z])', lambda m: m.group(1).upper(), s)

def lowerFirst(s):
	if len(s) > 2:
		return s[0].lower()  + s[1:]
	return s

if __name__=='__main__':
	items = parse()

	print "package mariamon"
	print 'import ('
	print '\t"errors"'
	print '\t"database/sql"'
	print ')'
	print "type pStatus struct {"
	seen = {}
	for item in items:
		if item['name'] in seen:
			continue
		print "\t", to_camelcase(item['name']), '  int64 `json:"%s"`' % item['name'].strip()
		seen[item['name']] = True
	print "}"
	print "type Status struct {"
	print "   pStatus"
	print "}"
	print "func (status Status) MarshalJSON() ([]byte, error) {"
	print "      return json.Marshal(status)"
	print "}"
	print "func (status Status) UnmarshalJSON(blob []byte) error {"
	print "      return json.Unmarshal(blob, &status.pStatus)"
	print "}"

	print
	seen = {}
	for item in items:
		if item['name'] in seen:
			continue
		print "// %s returns %s" % (to_camelcase(item['name']), lowerFirst(item['description']))
		print "func (s *Status) %s() int64 {" % to_camelcase(item['name'])
		print "\treturn s.pStatus.%s // %s" % (lowerFirst(to_camelcase(item['name'])), item['name'].strip())
		print "}"
		print
		seen[item['name']] = True

	print ' var ErrInvalidCast = errors.New("invalid cast")'
        print "func setInt64(src interface{}, dst *int64) error {"
        print "        val, ok := src.(int64)"
        print "        if !ok {"
        print "            return ErrInvalidCast"
        print "        }"
        print "        *dst = val"
        print "        return nil"
        print "}"
        print 
	print "func ReadStatus(db *sql.DB) (status Status, err error) {"
	print '        rows, err := db.Query("SHOW /*!50002 GLOBAL */ STATUS")'
        print '        if err != nil {'
	print '            return status, err'
	print '        }'
	print
	print '        for rows.Next() {'
	print '             var name string'
        print '             var val interface{}'
        print '             err = rows.Scan(&name, &val)'
        print '             if err != nil {'
        print '                  return status, err'
        print '             }'
        print '             switch name {'
	seen = {}
	for item in items:
		if item['name'] in seen:
			continue
	        print '	            case "%s":' % item['name']
	        print '	               err = setInt64(val, &status.pStatus.%s)' % to_camelcase(item['name'])
		seen[item['name']] = True
 	print '	            default:'
        print '	            }'
        print '             if err != nil {'
        print '                 return status, err'
        print '             }'
        print '        }'
	print '        return status, err'
        print '}'

