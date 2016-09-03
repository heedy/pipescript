var pipescript = (pipescript === undefined
    ? require("../pipescript.js")
    : pipescript);
var expect = (expect === undefined
    ? require("chai").expect
    : expect);

describe("PipeScript Javascript", function() {
    it("should give parse error on invalid script", function() {
        expect(function() {
            pipescript.Script("blah blah")
        }).to.throw();
    });
    it("should give no parse error on valid script", function() {
        expect(function() {
            pipescript.Script("$")
        }).to.not.throw()
    });
    it("should run identity script successfully", function() {
        expect(JSON.parse(pipescript.Script("$").Run('[{"t": 1, "d": 5},{"t": 2, "d": 6}'))).to.deep.equal([
            {
                t: 1.0,
                d: 5
            }, {
                t: 2.0,
                d: 6
            }
        ]);
        expect(pipescript.Script("$").Transform([
            {
                t: 1,
                d: 5
            }, {
                t: 2,
                d: 6
            }
        ])).to.deep.equal([
            {
                t: 1.0,
                d: 5
            }, {
                t: 2.0,
                d: 6
            }
        ]);
    });
    it("should run if script successfully", function() {
        expect(JSON.parse(pipescript.Script("if $ > 5").Run('[{"t": 1, "d": 5},{"t": 2, "d": 6}'))).to.deep.equal([
            {
                t: 2.0,
                d: 6
            }
        ]);
        expect(pipescript.Script("if $ > 5").Transform([
            {
                "t": 1,
                "d": 5
            }, {
                "t": 2,
                "d": 6
            }
        ])).to.deep.equal([
            {
                t: 2.0,
                d: 6
            }
        ]);
    });
    it("should have correct not prescendence", function() {
        expect(JSON.parse(pipescript.Script("$ < 0 or not $ < 1").Run('[{"t": 1, "d": 0.1}'))).to.deep.equal([
            {
                t: 1.0,
                d: false
            }
        ]);
    });
    it("should correctly handle activity example", function() {
        expect(JSON.parse(pipescript.Script('map($["activity"],$("steps"):sum)').Run('[{ \
                        "t": 1,\
                        "d": {\
                              "steps": 14,\
                              "activity": "walking"\
                          }\
                    },\
                    {\
                        "t": 2,\
                        "d": {\
                              "steps": 10,\
                              "activity": "running"\
                          }\
                    },\
                    {\
                        "t": 3,\
                        "d": {\
                              "steps": 12,\
                              "activity": "walking"\
                          }\
                    },\
                    {\
                        "t": 4,\
                        "d": {\
                              "steps": 5,\
                              "activity": "running"\
                          }\
                    }]'))).to.deep.equal([
            {
                t: 4.0,
                d: {
                    walking: 26,
                    running: 15
                }
            }
        ]);
    });
    it("should not have issues with local time zone", function() {
        expect(JSON.parse(pipescript.Script("yearmonth").Run('[{"t": 65000, "d": 5}]'))).to.deep.equal([
            {
                t: 65000.0,
                d: "January"
            }
        ]);
    });
    it("should be able to parse JSON", function() {
        expect(JSON.parse(pipescript.Script('map($["activity"],$("steps"):sum)').Run('[\
              {\
                "t": "1974-08-11T01:37:45+00:00",\
                "steps": 14,\
                "activity": "walking"\
            },\
              {\
                "t": "1974-08-12T03:44:25+00:00",\
                "steps": 10,\
                "activity": "running"\
            },\
              {\
                "t": "1974-08-12T04:17:45+00:00",\
                "steps": 12,\
                "activity": "walking"\
            },\
              {\
                "t": "1974-08-12T05:24:25+00:00",\
                "steps": 5,\
                "activity": "running"\
            }]', "json"))).to.deep.equal([
            {
                t: 145517065.0,
                d: {
                    walking: 26,
                    running: 15
                }
            }
        ]);
    });
    it("should be able to parse CSV", function() {
        expect(JSON.parse(pipescript.Script('map($["activity"],$("steps"):sum)').Run('time,steps,activity\n\
1974-08-11T01:37:45+00:00,14,walking\n\
1974-08-12T03:44:25+00:00,10,running\n\
1974-08-12T04:17:45+00:00, 12,walking\n\
1974-08-12T05:24:25+00:00,5,running', "csv"))).to.deep.equal([
            {
                t: 145517065.0,
                d: {
                    walking: 26,
                    running: 15
                }
            }
        ]);
    });
});
