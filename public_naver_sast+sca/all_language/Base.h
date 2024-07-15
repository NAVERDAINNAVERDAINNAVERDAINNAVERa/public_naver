#ifndef _BASE_H_
#define _BASE_H_

#include <string>

class Base
{
public:
	/* WRITE_VIRTUAL_DTOR_FOR_BASE_CLASS */
	Base() {}

	virtual std::string toString(char value);
	virtual std::string toString(char value, bool flag);
	virtual std::string toString(int value);
	virtual std::string toString(int value, bool flag);
	virtual std::string toString(float value);
	virtual std::string toString(float value, bool flag);
};

#endif /* _BASE_H_ */
