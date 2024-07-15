
#include "Base.h"

#include <stdio.h>
#include <algorithm>
#include <iostream>
#include <sstream>

using namespace std;

string Base::toString(char value)
{
	return this->toString(value, false);
}

string Base::toString(char value, bool flag)
{
	string ret;
	ret.resize( 1024 );
	if( flag == true )
	{
		cout << "message." << endl;
	}
	return itoa( value, &ret[0], 10 );
}

string Base::toString(int value)
{
	return this->toString(value, false);
}

string Base::toString(int value, bool flag)
{
	string ret;
	ret.resize( 1024 );
	if( flag == true )
	{
		cout << "message." << endl;
	}
	return itoa( value, &ret[0], 10 );

}

string Base::toString(float value)
{
	return this->toString(value, false);
}

string Base::toString(float value, bool flag)
{
	ostringstream o;
	o << value;
	if( flag == true )
	{
		cout << "message." << endl;
	}

	return o.str();
}


