var pipescript = (pipescript === undefined ? require("../pipescript.js") : pipescript);
var expect = (expect === undefined ? require("chai")
    .expect : expect);

describe("PipeScript Javascript", function() {
    it("should give parse error on invalid script", function() {
        expect(pipescript.Script("blah blah")
                .IsValid())
            .to.equal(false);
    });
    it("should give no parse error on valid script", function() {
        expect(pipescript.Script("$")
                .IsValid())
            .to.equal(true)
    });
    it("should run identity script successfully", function() {
        expect(JSON.parse(pipescript.Script("$")
                .Run('[{"t": 1, "d": 5},{"t": 2, "d": 6}')))
            .to.deep.equal([{
                t: 1.0,
                d: 5
            }, {
                t: 2.0,
                d: 6
            }]);
    });
    it("should run if script successfully", function() {
        expect(JSON.parse(pipescript.Script("if $ > 5")
                .Run('[{"t": 1, "d": 5},{"t": 2, "d": 6}')))
            .to.deep.equal([{
                t: 2.0,
                d: 6
            }]);
    });
    it("should correctly handle activity example", function() {
        expect(JSON.parse(pipescript.Script('map($["activity"],$("steps"):sum) | if last')
                .Run('[{ \
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
                    }]')))
            .to.deep.equal([{
                t: 4.0,
                d: {
                    walking: 26,
                    running: 15
                }
            }]);
    });
    it("should not have issues with local time zone", function() {
        expect(JSON.parse(pipescript.Script("yearmonth")
                .Run('[{"t": 65000, "d": 5}]')))
            .to.deep.equal([{
                t: 65000.0,
                d: "January"
            }]);
    });
})
